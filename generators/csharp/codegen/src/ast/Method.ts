import { Access } from "./Access";
import { Annotation } from "./Annotation";
import { ClassReference } from "./ClassReference";
import { CodeBlock } from "./CodeBlock";
import { AstNode } from "./core/AstNode";
import { DocXmlWriter } from "./core/DocXmlWriter";
import { Writer } from "./core/Writer";
import { Parameter } from "./Parameter";
import { Type } from "./Type";

export enum MethodType {
    INSTANCE,
    STATIC
}

export declare namespace Method {
    interface Args {
        /* The name of the method */
        name: string;
        /* The access of the method */
        access: Access;
        /* Whether the method is sync or async. Defaults to false. */
        isAsync?: boolean;
        /* The parameters of the method */
        parameters: Parameter[];
        /* Whether the method overrides a method in it's base class */
        override?: boolean;
        /* The return type of the method */
        return_?: Type;
        /* The body of the method */
        body?: CodeBlock;
        /* Summary for the method */
        summary?: string;
        /* The type of the method, defaults to INSTANCE */
        type?: MethodType;
        /* The class this method belongs to, if any */
        classReference?: ClassReference;
        /* Any annotations to add to the method */
        annotations?: Annotation[];
        /* Any code example to add to the method */
        codeExample?: string;
    }
}

export class Method extends AstNode {
    public readonly name: string;
    public readonly isAsync: boolean;
    public readonly access: Access;
    public readonly return: Type | undefined;
    public readonly body: CodeBlock | undefined;
    public readonly summary: string | undefined;
    public readonly type: MethodType;
    public readonly reference: ClassReference | undefined;
    public readonly override: boolean;
    private readonly parameters: Parameter[];
    private readonly annotations: Annotation[];
    private readonly codeExample: string | undefined;

    constructor({
        name,
        isAsync,
        override,
        access,
        return_,
        body,
        summary,
        type,
        classReference,
        parameters,
        annotations,
        codeExample
    }: Method.Args) {
        super();
        this.name = name;
        this.isAsync = isAsync ?? false;
        this.override = override ?? false;
        this.access = access;
        this.return = return_;
        this.body = body;
        this.summary = summary;
        this.type = type ?? MethodType.INSTANCE;
        this.reference = classReference;
        this.parameters = parameters;
        this.annotations = annotations ?? [];
        this.codeExample = codeExample;
    }

    public addParameter(parameter: Parameter): void {
        this.parameters.push(parameter);
    }

    public write(writer: Writer): void {
        const docXmlWriter = new DocXmlWriter(writer);
        if (this.summary != null) {
            docXmlWriter.writeNodeWithEscaping("summary", this.summary);
        }
        if (this.codeExample != null) {
            docXmlWriter.writeOpenNode("example");
            docXmlWriter.writeOpenNode("code");
            docXmlWriter.writeMultilineWithEscaping(this.codeExample);
            docXmlWriter.writeCloseNode("code");
            docXmlWriter.writeCloseNode("example");
        }

        if (this.annotations.length > 0) {
            writer.write("[");
            this.annotations.forEach((annotation) => {
                annotation.write(writer);
            });
            writer.write("]");
            writer.writeNewLineIfLastLineNot();
        }

        writer.write(`${this.access} `);
        if (this.type === MethodType.STATIC) {
            writer.write("static ");
        }
        if (this.isAsync) {
            writer.write("async ");
        }
        if (this.override) {
            writer.write("override ");
        }
        if (this.return == null) {
            if (this.isAsync) {
                writer.writeNode(
                    new ClassReference({
                        name: "Task",
                        namespace: "System.Threading.Tasks"
                    })
                );
                writer.write(" ");
            } else {
                writer.write("void ");
            }
        } else {
            if (this.isAsync) {
                writer.write("Task<");
                this.return.write(writer);
                writer.write(">");
            } else {
                this.return.write(writer);
            }
            writer.write(" ");
        }
        writer.write(`${this.name}(`);
        this.parameters.forEach((parameter, idx) => {
            parameter.write(writer);
            if (idx < this.parameters.length - 1) {
                writer.write(", ");
            }
        });
        writer.write(")");
        writer.writeLine(" {");

        writer.indent();
        this.body?.write(writer);
        writer.dedent();

        writer.writeLine("}");
    }

    public getParameters(): Parameter[] {
        return this.parameters;
    }
}
